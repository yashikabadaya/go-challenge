Initial solution
-----------------

Iteratively query the server for each url query. Perform merge sort on the arrays returned (considering to remove duplicates while sorting).
Worst case time complexity : O(nlogn)

Other Ideas
------------

If we assume the urls to always return same results then we can cache those results in sorted manner. Whenever the user queries from the api, majority of time will be spent only in sorting and removing duplicates from the cached results. The time spent to query from the cache is considered negligible compared to sorting and removal of duplicates.

 
