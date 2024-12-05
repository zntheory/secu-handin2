using Microsoft.AspNetCore.Mvc;
namespace hospital;

[ApiController]
[Route("[controller]")]
public class Controller : ControllerBase
{
    private static Hospital _hospital;

    public Controller(Hospital hospital)
    {
        _hospital = hospital;
    }

    [HttpPost("Aggregate")]
    public IActionResult Post([FromBody] int input)
    {
        if (_hospital.ShareCount >= 3) return BadRequest(new { message = "No more shares, pwease!!" });
        _hospital.Aggregate(input);
        return Ok(new { message = $"Received: {input}" });
    }
}