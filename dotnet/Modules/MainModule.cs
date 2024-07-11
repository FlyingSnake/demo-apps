using Carter;
using Dapper;
using MySqlConnector;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace demoApp.Modules
{
  public class MainModule : CarterModule
  {
    public class User
    {
      public int Id { get; set; }
      public string Name { get; set; }
      public string Email { get; set; }
    }

    public MainModule()
    {
      Get("/", async (req, res) => await res.WriteAsync("Hello World (.NET / Carter)"));

      Get("/users", async (req, res) =>
      {
        var connection = req.HttpContext.RequestServices.GetService<MySqlConnection>();
        var users = await connection.QueryAsync<User>("SELECT * FROM user");
        await res.WriteAsJsonAsync(users);
      });

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
        var random = new Random();
        var (statusCode, message) = statusList[random.Next(statusList.Count)];
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
