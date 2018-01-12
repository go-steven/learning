<?php
// download apache-log4php-2.3.0 to log4php directory.
// run it with: 
//    php -f index.php

require_once dirname(__FILE__).'/log4php/Logger.php';

Logger::configure('log4php.xml');

$logger = Logger::getLogger('test');
$logger->debug('Hello again!'); 