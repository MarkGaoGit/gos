<?php

class JsonRPC {
    private $conn;

    public function __construct($host, $port)
    {
        $this->conn = fsockopen($host, $port);
        if (!$this->conn) {
            return false;
        }
    }

    public function Call($method, $params)
    {
        // 这里发现 传入的params外层必须需要用数组包含，
        // 最外层需要把发送的真实数据使用数组包起来
        $err = fwrite($this->conn, json_encode([
            'method' => $method,
            'params' => [
                $params
            ],
            'id' => rand(1, 100)
        ]) . "\n");

        if ($err === false) {
            return false;
        }

        stream_set_timeout($this->conn, 0, 3000);
        $line = fgets($this->conn);
        if ($line === false) {
            return '没数据';
        }

        return json_decode($line, true);
    }
}


//使用PHP调用GO提供的RPC服务
$user = [
    'name' => 'Michael',
    'age' => 30,
];
$rpcServer = new JsonRPC('10.2.4.216', '2000');
$response = $rpcServer->Call('userService.UserCreate', $user);
var_dump($response);

