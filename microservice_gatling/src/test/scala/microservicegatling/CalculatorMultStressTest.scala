package microservicegatling

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class CalculatorMultStressTest extends Simulation {

  val config = new Config("http://localhost:8080")

  val scn = scenario("History Multiplication Calculator Stress Test")
    .exec(http("First mult")
      .get("/calc/sub/8/8"))
    .exec(http("Second mult")
      .get("/calc/mult/7/5"))
    .exec(http("Third mult")
      .get("/calc/mult/10/5"))
    .pause(1000 milliseconds)
    .exec(http("History mult")
      .get("/calc/history"))

  setUp(
    scn.inject(constantUsersPerSec(config.quantityUsers).during(config.durationStressTest)
  ).protocols(config.getHttpConfig))
}
