package microservicegatling

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class CalculatorSubStressTest extends Simulation {

  val config = new Config("http://localhost:8080")

  val scn = scenario("History Subtraction Calculator Stress Test")
    .exec(http("First sub")
      .get("/calc/sub/8/8"))
    .exec(http("Second sub")
      .get("/calc/sub/7/5"))
    .exec(http("Third sub")
      .get("/calc/sub/10/5"))
    .pause(1000 milliseconds)
    .exec(http("History sub")
      .get("/calc/history"))

  setUp(
    scn.inject(constantUsersPerSec(config.quantityUsers).during(config.durationStressTest)
  ).protocols(config.getHttpConfig))
}
