package microservicegatling

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class CalculatorSumStressTest extends Simulation {

  val config = new Config("http://localhost:8080")

  val scn = scenario("History Sum Calculator Stress Test")
    .exec(http("First sum")
      .get("/calc/sum/8/8"))
    .exec(http("Second sum")
      .get("/calc/sum/7/5"))
    .exec(http("Third sum")
      .get("/calc/sum/10/5"))
    .pause(1000 milliseconds)
    .exec(http("History sum")
      .get("/calc/history"))

  setUp(
    scn.inject(constantUsersPerSec(config.quantityUsers).during(config.durationStressTest)
  ).protocols(config.getHttpConfig))
}
