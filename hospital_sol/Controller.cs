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

    [HttpPost("Shares")]
    public IActionResult Post([FromBody] int result)
    {
        if (_hospital.ShareCount <= 3)
        {
            _hospital.Aggregate(result);
            return Ok(new{message = $"Received result: {result}"});
        }

        return BadRequest(new {message = "Stawp invading secrets pwease!!"});
    }
}
