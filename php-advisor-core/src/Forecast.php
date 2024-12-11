<?php

namespace StormGeo\AdvisorCore;

/**
 * @package StormGeo\AdvisorCore
 */
class Forecast extends BaseRouter
{
  /**
   * @param   array $payload
   * @return  array
   */
  public function getDaily($payload)
  {
    return parent::makeRequest('GET', '/v1/forecast/daily', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getHourly($payload)
  {
    return parent::makeRequest('GET', '/v1/forecast/hourly', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getPeriod($payload)
  {
    return parent::makeRequest('GET', '/v1/forecast/period', $payload);
  }
}
