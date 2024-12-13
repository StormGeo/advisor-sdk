<?php

namespace StormGeo\AdvisorCore\Payloads;

class StationPayload
{
  /**
   * @var string
   */
  public $stationId;

  /**
   * @var string
   */
	public $startDate;

  /**
   * @var string
   */
  public $endDate;

  /**
   * @var array<string>
   */
	public $variables;

  /**
   * @var string
   */
	public $layer;

  /**
   * @var array
   */
  public function getQueryParams()
  {
    return (array) $this;
  }
}
