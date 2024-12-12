<?php

namespace StormGeo\AdvisorCore\Routes;

/**
 * @package StormGeo\AdvisorCore
 */
class Chart extends BaseRouter
{
  /**
   * @param   array $payload
   * @return  array
   */
  public function getForecastDaily($payload)
  {
    return parent::makeRequest('GET', '/v1/forecast/daily/chart', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getForecastHourly($payload)
  {
    return parent::makeRequest('GET', '/v1/forecast/hourly/chart', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getObservedDaily($payload)
  {
    return parent::makeRequest('GET', '/v1/observed/daily/chart', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getObservedHourly($payload)
  {
    return parent::makeRequest('<?php

namespace StormGeo\AdvisorCore;GET', '/v1/observed/hourly/chart', $payload);
  }
}
