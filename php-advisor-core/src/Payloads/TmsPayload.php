<?php

namespace StormGeo\AdvisorCore\Payloads;

class TmsPayload
{
  /**
   * @var string
   */
  public $istep;

  /**
   * @var string
   */
	public $fstep;

  /**
   * @var string
   */
	public $server;

  /**
   * @var string
   */
	public $mode;

  /**
   * @var string
   */
	public $variable;

  /**
   * @var string
   */
	public $aggregation;

  /**
   * @var int
   */
	public $x;

  /**
   * @var int
   */
	public $y;

  /**
   * @var int
   */
	public $z;

  /**
   * @return array
   */
  public function getQueryParams()
  {
    return [
      'istep' => $this->istep,
			'fstep' => $this->fstep,
    ];
  }
}
