package microservicegatling

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class CalculatorDivStressTest extends Simulation {

  val config = new Config("http://localhost:8080")

  val scn = scenario("History Division Calculator Stress Test")
    .exec(http("First div")
      .get("/calc/div/8/8"))
    .exec(http("Second div")
      .get("/calc/div/7/5"))
    .exec(http("Third div")
      .get("/calc/div/10/5"))
    .pause(1000 milliseconds)
    .exec(http("History div")
      .get("/calc/history"))

  setUp(
    scn.inject(constantUsersPerSec(config.quantityUsers).during(config.durationStressTest)
  ).protocols(config.getHttpConfig))
}
