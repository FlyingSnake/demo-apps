using Carter;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using System;
using MySqlConnector;

namespace demoApp
{
  public class Startup
  {
    public Startup(IConfiguration configuration)
    {
      Configuration = configuration;
    }

    public IConfiguration Configuration { get; }

    public void ConfigureServices(IServiceCollection services)
    {
      services.AddCarter();
      services.AddTransient<MySqlConnection>(_ =>
      {
        var connectionString = new MySqlConnectionStringBuilder
        {
          Server = Environment.GetEnvironmentVariable("DB_HOST"),
          UserID = Environment.GetEnvironmentVariable("DB_USERNAME"),
          Password = Environment.GetEnvironmentVariable("DB_PASSWORD"),
          Database = Environment.GetEnvironmentVariable("DB_DATABASE")
        }.ConnectionString;

        return new MySqlConnection(connectionString);
      });
    }

    public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
    {
      if (env.IsDevelopment())
      {
        app.UseDeveloperExceptionPage();
      }

      app.UseRouting();
      app.UseEndpoints(builder => builder.MapCarter());
    }
  }
}
