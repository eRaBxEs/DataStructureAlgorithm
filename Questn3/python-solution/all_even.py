def all_even(nums):
  is_even = False
  even_count = 0
  odd_count = 0
  for num in nums:
    if num%2==0 :
      even_count = even_count + 1
    else:
      odd_count = odd_count + 1

  if even_count == len(nums):
    is_even = True
  else:
    is_even = False

  return is_even