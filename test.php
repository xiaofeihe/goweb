<?php
$begin = microtime();

for($i= 0; $i < 6; $i++) {
 test($i==5);
}

echo howLong($begin, microtime()); //输出页面执行微毫秒数
function howLong($begin, $end){
 $begin = explode(' ',$begin);
 $end = explode(' ',$end);
 return $end = ($end[1]-$begin[1])+($end[0]-$begin[0]);
}

function test($c){
$x=0;
for ($i= 0; $i < 1000000; $i++) {
$x += $i;
for($m = 0; $m < 50; $m++ ){
$x += 0;
}
}
if($c)
print_r('OK:耗时：');

}

?>
