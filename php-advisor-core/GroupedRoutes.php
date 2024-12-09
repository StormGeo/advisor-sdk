<?php

namespace StormGeo\AdvisorCore;

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
    return parent::makeRequest('GET', '/v1/observed/hourly/chart', $payload);
  }
}

/**
 * @package StormGeo\AdvisorCore
 */
class Climatology extends BaseRouter
{
  /**
   * @param   array $payload
   * @return  array
   */
  public function getDaily($payload)
  {
    return parent::makeRequest('GET', '/v1/climatology/daily', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getMonthly($payload)
  {
    return parent::makeRequest('GET', '/v1/climatology/monthly', $payload);
  }
}

/**
 * @package StormGeo\AdvisorCore
 */
class CurrentWeather extends BaseRouter
{
  /**
   * @param   array $payload
   * @return  array
   */
  public function get($payload)
  {
    return parent::makeRequest('GET', '/v1/current-weather', $payload);
  }
}

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

/**
 * @package StormGeo\AdvisorCore
 */
class Monitoring extends BaseRouter
{
  /**
   * @return  array
   */
  public function getAlerts()
  {
    return parent::makeRequest('GET', '/v1/monitoring/alerts');
  }
}

/**
 * @package StormGeo\AdvisorCore
 */
class Observed extends BaseRouter
{
  /**
   * @param   array $payload
   * @return  array
   */
  public function getDaily($payload)
  {
    return parent::makeRequest('GET', '/v1/observed/daily', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getHourly($payload)
  {
    return parent::makeRequest('GET', '/v1/observed/hourly', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getPeriod($payload)
  {
    return parent::makeRequest('GET', '/v1/observed/period', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getLightning($payload)
  {
    return parent::makeRequest('GET', '/v1/observed/lightning', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function postLightning($payload)
  {
    return parent::makeRequest('POST', '/v1/observed/lightning', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getFireFocus($payload)
  {
    return parent::makeRequest('GET', '/v1/observed/fire-focus', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function postFireFocus($payload)
  {
    return parent::makeRequest('POST', '/v1/observed/fire-focus', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getStationData($payload)
  {
    return parent::makeRequest('GET', '/v1/station', $payload);
  }
}

/**
 * @package StormGeo\AdvisorCore
 */
class Plan extends BaseRouter
{
  /**
   * @return  array
   */
  public function getInfo()
  {
    return parent::makeRequest('GET', '/v1/plan/:token');
  }
}

/**
 * @package StormGeo\AdvisorCore
 */
class Schema extends BaseRouter
{
  /**
   * @return  array
   */
  public function getDefinition()
  {
    return parent::makeRequest('GET', '/v1/schema/definition');
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function postDefinition($payload)
  {
    return parent::makeRequest('POST', '/v1/schema/definition', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function postParameters($payload)
  {
    return parent::makeRequest('POST', '/v1/schema/parameters', $payload);
  }
}

/**
 * @package StormGeo\AdvisorCore
 */
class Tms extends BaseRouter
{
  /**
   * @param   array $payload
   * @return  array
   */
  public function get($payload)
  {
    return parent::makeRequest('GET', '/v1/tms');
  }
}
