<?php

namespace StormGeo\AdvisorCore;

use StormGeo\AdvisorCore\Routes\Chart;
use StormGeo\AdvisorCore\Routes\Climatology;
use StormGeo\AdvisorCore\Routes\CurrentWeather;
use StormGeo\AdvisorCore\Routes\Forecast;
use StormGeo\AdvisorCore\Routes\Monitoring;
use StormGeo\AdvisorCore\Routes\Observed;
use StormGeo\AdvisorCore\Routes\Plan;
use StormGeo\AdvisorCore\Routes\Schema;
use StormGeo\AdvisorCore\Routes\Tms;

/**
 * @package StormGeo\AdvisorCore
 */
class AdvisorCore
{
  /**
   * @var Chart
   */
  public $chart;

  /**
   * @var Climatology
   */
  public $climatology;

  /**
   * @var CurrentWeather
   */
  public $currentWeather;

  /**
   * @var Forecast
   */
  public $forecast;

  /**
   * @var Monitoring
   */
  public $monitoring;

  /**
   * @var Observed
   */
  public $observed;

  /**
   * @var Plan
   */
  public $plan;

  /**
   * @var Schema
   */
  public $schema;

  /**
   * @var Tms
   */
  public $tms;

  /**
   * @param   string  $token
   * @param   int     $retries
   * @param   int     $delay
   */
  public function __construct($token, $retries = 5, $delay = 5)
  {
    $this->chart = new Chart($token, $retries, $delay);
    $this->climatology = new Climatology($token, $retries, $delay);
    $this->currentWeather = new CurrentWeather($token, $retries, $delay);
    $this->forecast = new Forecast($token, $retries, $delay);
    $this->monitoring = new Monitoring($token, $retries, $delay);
    $this->observed = new Observed($token, $retries, $delay);
    $this->plan = new Plan($token, $retries, $delay);
    $this->schema = new Schema($token, $retries, $delay);
    $this->tms = new Tms($token, $retries, $delay);
  }
}
