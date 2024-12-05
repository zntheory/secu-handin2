using Microsoft.AspNetCore.Mvc;
namespace patient;

[ApiController]
[Route("[controller]")]
public class PatientController : ControllerBase
{
    private readonly Patient _patient;

    public PatientController(Patient patient)
    {
        _patient = patient;
    }

    [HttpPost("Message")]
    public ActionResult<string> Post([FromBody] int input)
    {
        _patient.AppendShare(input);
        return Ok(new {message = $"Fellow patient received share: {input}"});
    }

}