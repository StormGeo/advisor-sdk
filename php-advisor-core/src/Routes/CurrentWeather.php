<?php

namespace StormGeo\AdvisorCore\Routes;

use StormGeo\AdvisorCore\Payloads\CurrentWeatherPayload;

/**
 * @package StormGeo\AdvisorCore
 */
class CurrentWeather extends BaseRouter
{
  /**
   * @param   CurrentWeatherPayload $payload
   * @return  array
   */
  public function get($payload)
  {
    return parent::makeRequest(
      'GET',
      '/v1/current-weather' . $this->formatQueryParams($payload->getQueryParams())
    );
  }
}
