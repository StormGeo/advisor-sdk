<?php

namespace StormGeo\AdvisorCore\Routes;

use StormGeo\AdvisorCore\Payloads\ClimatologyPayload;

/**
 * @package StormGeo\AdvisorCore
 */
class Climatology extends BaseRouter
{
  /**
   * @param   ClimatologyPayload $payload
   * @return  array
   */
  public function getDaily($payload)
  {
    return parent::makeRequest(
      'GET',
      '/v1/climatology/daily' . $this->formatQueryParams($payload->getQueryParams())
    );
  }

  /**
   * @param   ClimatologyPayload $payload
   * @return  array
   */
  public function getMonthly($payload)
  {
    return parent::makeRequest(
      'GET',
      '/v1/climatology/monthly' . $this->formatQueryParams($payload->getQueryParams())
    );
  }
}
