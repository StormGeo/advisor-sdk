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
   * @param   string  $token      API Token
   * @param   int     $attempts   Number of attempts if an error occurs
   * @param   int     $delay      Delay between attempts
   */
  public function __construct($token, $attempts = 5, $delay = 5)
  {
    $this->chart = new Chart($token, $attempts, $delay);
    $this->climatology = new Climatology($token, $attempts, $delay);
    $this->currentWeather = new CurrentWeather($token, $attempts, $delay);
    $this->forecast = new Forecast($token, $attempts, $delay);
    $this->monitoring = new Monitoring($token, $attempts, $delay);
    $this->observed = new Observed($token, $attempts, $delay);
    $this->plan = new Plan($token, $attempts, $delay);
    $this->schema = new Schema($token, $attempts, $delay);
    $this->tms = new Tms($token, $attempts, $delay);
  }
}
