<?php

require_once __DIR__ .'/../vendor/autoload.php';

define("APP_ROOT", __DIR__ .'/..');

$class = $argv[1];
$func = $argv[2];
$msg = $argv[3];

call_user_func_array([$class, $func], ['msg' => $msg]);