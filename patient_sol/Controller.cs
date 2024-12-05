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
    public ActionResult<string> Message([FromBody] int input)
    {
        _patient.AppendShare(input);
        return Ok(new {message = $"Received share: {input}"});
    }

}