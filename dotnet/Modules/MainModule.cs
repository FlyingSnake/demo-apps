using Carter;
using Microsoft.AspNetCore.Http;
using MySql.Data.MySqlClient;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace demoApp.Modules
{
  public class MainModule : CarterModule
  {
    private readonly string connectionString;

    public MainModule()
    {
      var dbHost = Environment.GetEnvironmentVariable("DB_HOST");
      var dbUsername = Environment.GetEnvironmentVariable("DB_USERNAME");
      var dbPassword = Environment.GetEnvironmentVariable("DB_PASSWORD");
      var dbName = Environment.GetEnvironmentVariable("DB_DATABASE");

      connectionString = $"Server={dbHost};Database={dbName};User={dbUsername};Password={dbPassword};";

      Get("/", async (req, res) => await res.WriteAsync("Hello World (.NET / Carter)"));

      Get("/status/random", async (req, res) =>
      {
        var statusList = new List<(int, string)>
          {
                    (200, "OK"),
                    (201, "Created"),
                    (202, "Accepted"),
                    (204, "No Content"),
                    (400, "Bad Request"),
                    (401, "Unauthorized"),
                    (403, "Forbidden"),
                    (404, "Not Found"),
                    (500, "Internal Server Error"),
                    (501, "Not Implemented"),
                    (502, "Bad Gateway"),
                    (503, "Service Unavailable")
          };
        var (statusCode, message) = statusList[new Random().Next(statusList.Count)];
        res.StatusCode = statusCode;
        await res.WriteAsJsonAsync(message);
      });

      Get("/sleep/{seconds:int}", async (req, res) =>
      {
        if (int.TryParse(req.RouteValues["seconds"].ToString(), out int seconds))
        {
          await Task.Delay(seconds * 1000);
          await res.WriteAsync($"sleep {seconds}s");
        }
        else
        {
          res.StatusCode = 400;
          await res.WriteAsync("Invalid input for seconds.");
        }
      });

      Get("/users", async (req, res) =>
      {
        var users = new List<object>();
        try
        {
          using (var connection = new MySqlConnection(connectionString))
          {
            await connection.OpenAsync();
            using (var command = new MySqlCommand("SELECT * FROM user", connection))
            using (var reader = await command.ExecuteReaderAsync())
            {
              while (await reader.ReadAsync())
              {
                users.Add(new
                {
                  id = reader["id"],
                  name = reader["name"],
                  email = reader["email"]
                });
              }
            }
          }
        }
        catch (Exception ex)
        {
          res.StatusCode = 500;
          await res.WriteAsJsonAsync(new { error = ex.Message });
          return;
        }
        await res.WriteAsJsonAsync(users);
      });

      Get("/exception", async (req, res) =>
      {
        var currentTime = DateTime.Now.ToString("yyyy-MM-dd HH:mm:ss");
        var errorMessage = "Internal Server Error - Manual Exception";
        Console.Error.WriteLine(errorMessage);
        res.StatusCode = 500;
        await res.WriteAsJsonAsync(new { timestamp = currentTime, status = 500, error = errorMessage });
      });
    }
  }
}
