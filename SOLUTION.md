Initial solution
-----------------

Created a go routine to iteratively query the server for each url query, append response to result arr and then perform merge sort on the result arr (considering to remove duplicates while sorting). In the foreground, there is wait time for 500ms and then function exists.
Worst case time complexity : O(nlogn)
Since it is mentioned to return an empty list as result only if all URLs returned errors or took too long to respond and the timeout must be respected regardless of the size of the data, consistency is preferred over availabilty in this solution. Importance is given to giving correct answers or return empty array if time limit exceeds.  

Other Ideas
------------

If we assume the urls to always return same results then we can cache those results in sorted manner. Whenever the user queries from the api, majority of time will be spent only in sorting and removing duplicates from the cached results. The time spent to query from the cache is considered negligible compared to sorting and removal of duplicates.

 
