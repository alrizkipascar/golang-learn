const myMap = new Map([
  ["1", "a String key"],
  [1, "a number key"],
  [true, true],
]);

console.log(myMap);

/* output
  Map(3) { '1' => 'a String key', 1 => 'a number key', true => true }
  */

const wrongMap = new Map();
wrongMap["My Key"] = "My Value";

console.log(wrongMap.has("My Key"));
console.log(wrongMap.delete("My Key"));

/* output
  false
  false
   */

const numberSet = new Set([1, 4, 6, 4, 1]);
numberSet.add(5);
numberSet.add(10);
numberSet.add(6);

numberSet.delete(4);

console.log(numberSet);

/* output
Set(4) { 1, 6, 5, 10 }
*/

const visitsCountMap = new Map(); // Menyimpan daftar user

function countUser(user) {
  let count = visitsCountMap.get(user) || 0;
  visitsCountMap.set(user, count + 1);
}

let jonas = { name: "Jonas" };
countUser(jonas); // Menambahkan user "Jonas"

jonas = null; // Data object "Jonas" dihapus

// delay dibutuhkan untuk menunggu garbage collector bekerja
setTimeout(function () {
  console.log(visitsCountMap);
  let count = visitsCountMap.get(jonas);

  console.log("count", count);
}, 1000);

/* output
Map(1) { { name: 'Jonas' } => 1 }
*/

const { inspect } = require("util");

const visitsCountMap = new WeakMap(); // Menyimpan daftar user

function countUser(user) {
  let count = visitsCountMap.get(user) || 0;
  visitsCountMap.set(user, count + 1);
}

let jonas = { name: "Jonas" };
countUser(jonas); // Menambahkan user "Jonas"

jonas = null; // Data object "Jonas" dihapus

// delay dibutuhkan untuk menunggu garbage collector bekerja
setTimeout(function () {
  console.log(inspect(visitsCountMap, { showHidden: true }));
}, 10000);

/* output
  WeakMap {  }
*/

const { inspect } = require("util");

const visitsCountMap = new WeakMap(); // Menyimpan daftar user

function countUser(user) {
  let count = visitsCountMap.get(user) || 0;
  visitsCountMap.set(user, count + 1);
}

let jonas = { name: "Jonas" };
countUser(jonas); // Menambahkan user "Jonas"

jonas = null; // Data object "Jonas" dihapus

// delay dibutuhkan untuk menunggu garbage collector bekerja
setTimeout(function () {
  console.log(inspect(visitsCountMap, { showHidden: true }));
}, 10000);

/* output
  WeakMap {  }
*/
