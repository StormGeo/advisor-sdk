<?php

namespace StormGeo\AdvisorCore\Routes;

/**
 * @package StormGeo\AdvisorCore
 */
abstract class BaseRouter
{
  const BASE_URL = 'http://advisor-core.climatempo.io/api';

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
   * @param   string            $method
   * @param   string            $route
   * @param   array             $body
   * @return  array|string|null
   */
  protected function makeRequest($method, $route, $body = [])
  {
    if ($method === 'GET' || $method === 'GET_IMAGE') {
      return $this->retryRequest(
        function() use ($method, $route) {
          return $this->makeGetRequest($this::BASE_URL . $route, $method === 'GET_IMAGE');
        },
        $this->retries,
        $this->delay
      );
    }
  }

  /**
   * @param   string      $url
   * @param   bool        $binaryReturn
   * @return  array
   */
  protected function makeGetRequest($url, $binaryReturn = false)
  {
    $ch = curl_init($url);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_HTTPGET, true);

    $response = curl_exec($ch);
    $responseInfo = curl_getinfo($ch, CURLINFO_HTTP_CODE);

    if ($response != false) {
      return [
        'statusCode' => $responseInfo,
        'data' => ($binaryReturn) ? $response : json_decode($response, true)
      ];
    }

    return [
      'statusCode' => null,
      'data' => null
    ];
  }

  /**
   * @param   callable(): (array|null)  $request
   * @param   int                       $retries
   * @param   int                       $delay
   * @return  array|null
   */
  protected function retryRequest($request, $retries, $delay)
  {
    $data = null;

    for ($retryNumber = $retries; $retryNumber >= 0; $retryNumber--) {
      $response = $request();
      $status = $response['statusCode'];
      $data = $response['data'];

      if (!is_null($status) && $status < 500 && $status != 429) {
        return $data;
      }

      if ($retryNumber > 0) {
        sleep($delay);
      }
    }

    return $data;
  }

  /**
   * @param   array   $queryParams
   * @return  string
   */
  protected function formatQueryParams($queryParams)
  {
    $params = "?token={$this->token}";
    $formattedParams = http_build_query($queryParams);
    if (strlen($formattedParams) > 0) {
      $params .= "&{$formattedParams}";
      $params = preg_replace('/%5B[0-9]+%5D/simU', '[]', $params);
    }

    return $params;
  }
}
