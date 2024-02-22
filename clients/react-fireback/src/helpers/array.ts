export function move(array: Array<any>, from: number, to: number) {
  if (to === from) return array;

  var target = array[from];
  var increment = to < from ? -1 : 1;

  for (var k = from; k != to; k += increment) {
    array[k] = array[k + increment];
  }
  array[to] = target;
  return array;
}
