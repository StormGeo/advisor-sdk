<?php

namespace StormGeo\AdvisorCore\Routes;

use StormGeo\AdvisorCore\Payloads\BasePayload;

/**
 * @package StormGeo\AdvisorCore
 */
class Forecast extends BaseRouter
{
  /**
   * @param   BasePayload $payload
   * @return  array
   */
  public function getDaily($payload)
  {
    return parent::makeRequest(
      'GET',
      '/v1/forecast/daily' . $this->formatQueryParams($payload->getQueryParams())
    );
  }

  /**
   * @param   BasePayload $payload
   * @return  array
   */
  public function getHourly($payload)
  {
    return parent::makeRequest(
      'GET',
      '/v1/forecast/hourly' . $this->formatQueryParams($payload->getQueryParams())
    );
  }

  /**
   * @param   BasePayload $payload
   * @return  array
   */
  public function getPeriod($payload)
  {
    return parent::makeRequest(
      'GET',
      '/v1/forecast/period' . $this->formatQueryParams($payload->getQueryParams())
    );
  }
}
