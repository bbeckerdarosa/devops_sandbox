package microservicegatling

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._
import io.gatling.core.protocol.Protocol

class Config(var baseUrl: String) {

  def getHttpConfig() : Protocol = {
    return http
            .baseUrl(this.baseUrl)
            .acceptHeader("text/html,application/xhtml+xml,application/xml;q=0.9,/;q=0.8")
            .acceptLanguageHeader("en-US,en;q=0.5")
            .userAgentHeader("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0")
  }

  def quantityUsers() : Int = {
    return Integer.getInteger("users", 1).toInt
  }

  def durationStressTest() : Int = {
    return Integer.getInteger("duration", 60).toInt
  }
}



