<?php

namespace StormGeo\AdvisorCore;

/**
 * @package StormGeo\AdvisorCore
 */
abstract class BaseRouter
{
  /**
   * @var string
   */
  protected $token;

  /**
   * @var int
   */
  protected $retries;

  /**
   * @var int
   */
  protected $delay;

  /**
   * @param   string  $token
   * @param   int     $retries
   * @param   int     $delay
   */
  public function __construct($token, $retries, $delay)
  {
    $this->token = $token;
    $this->retries = $retries;
    $this->delay = $delay;
  }

  /**
   * @param   string  $method
   * @param   string  $route
   * @param   array   $payload
   * @return  array
   */
  protected function makeRequest($method, $route, $payload)
  {}
}
