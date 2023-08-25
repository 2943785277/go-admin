var search = function(nums, target) {
  const len = nums.length;
  let left = 0;
  let right = len - 1;

  while(left <= right) {
    console.log(left,'----',right)
    const mid = (left + right) >>> 1;
    console.log(nums[mid],'----',mid)
    console.log('--------------------')
    if (target === nums[mid]) {
      return mid;
    } else if (target < nums[mid]) {
      right = mid - 1;
    } else {
      left = mid + 1;
    }
  }

  return -1;
};

var search1 = function(nums, target) {
  for (let index = 0; index < nums.length; index++) {
    if(target == nums[index]) {
      return nums[index]
    }
  }
}

console.log(search([1,3,3,4,5,8,9,11,12],8))


