 /**
 * @param $departure string
 * @param $arrival string
 * @param $places array
 * @return array
 */

function findRoute($departure, $arrival, $places, $fare = 0, $route = [], $result = [[], 0])
{
    foreach ($places as $path) {
        if ($path[0] == $departure) {
            if ($path[1] == $arrival) {
                if ($result[1] > $fare || !$result[1]) {
                    $result = [array_merge($route, [$path]), $fare + $path[2]];
                    break;
                }
            } else {
                $result = findRoute($path[1], $arrival, $places, $fare + $path[2], array_merge($route, [$path]), $result);
            }
        }
    }
    return $result;
}

/**
 * /*$places = [["A", "B", 1],
 * ["B", "C", 5],
 * ["C", "D", 2],
 * ["B", "D", 1],
 * ["A", "C", 2]
 * ];
 * print_r(findRoute("A", "D", $places));
 */
