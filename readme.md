# MKPIS

Measuring the development process for Gitflow managed projects.


## How does it work?

Part of healthy engineering coulture is foster team and developers to continuously improv their work an processes.

In order to improve the Engeniering Managers/Software Managers/Tech leads need visibility over the development process.


Ofcourse there are tons of metrics you can extract from the whole development process, since the developer starts writing the code to the moment that code reaches production.

In many companies they realy in the agile defintion of Team velocity or 

But this project put specific focus on the feature pull request and the release pull request.



After deploy, there is lots of work too. Operations usually take over the risks. In many companies, sysadmins or DevOps perform rollback in production. Usually, it is not at developers’ hands too. The Product Team also measures the impact of the deployed changes to feed the backlog cyclically.

That’s why I like to visualize this part of the process separately. It gives more visibility to the development team and isolates the development from the full system.


## Visibility over the process



## Engineering Metrics 



### Feature Pull Request

<pre>
+----------+              +------+               +--------------+                 +-------------+               +-------+
|          |    Opens     |      |  Waits for    |              |  Discuss until  |             |   Waits for   |       |
|  Commit  +-------------->  PR  +---------------> First Review +-----------------> Last Review +---------------> Merge |
|          |              |      |               |              |                 |             |               |       |
+----------+              +------+               +--------------+                 +-------------+               +-------+

+                                              Feature Lead Time                                                        + 
+-----------------------------------------------------------------------------------------------------------------------+
+                                                                                                                       +
                                                               Pull Request Lead time
                          +                                                                                             +
                          +---------------------------------------------------------------------------------------------+
                          +                                                                                             +
                                                                 Review time
                                                 +                                              +
                                                 +----------------------------------------------+
                                                 +                                              +
                            Time to first review                                           Last review to merge time
                          +                      +                                  +                                   +
                          +----------------------+                                  +-----------------------------------+
                          +                      +                                  +                                   +

</pre>



**Feature Lead Time:** it measures how much time the first commit in the pull request takes to reach the devel branch.

`Formula: (merged_at  - first_commit_created_at)`

**Pull Request Lead time:** it mesures the time the code review process plus the corrections needed to move merge the code. It considers all the steps of code review and the wating times.

`Formula: (merged_at – opened_at)`

**Review time:** it mesures the time from the first to the last review in the pull request.
`Formula: (last_review_created_at  – first_review_created_at )`

**Time to First Review:** it mesures how much time the team takes to make the first code review. 

`Formula: (first_review_created_at - last_commit_created_at)`

**Last Review to Merge Time:** it 

`Formula: (merged_at - last_commit_created_at)`

**Pull request Size:** it measures the pull request size in terms of changes it contains.


`Formula: (pull_request_additions + pull_request_deletions)`


### Release Pull Request

<pre>
+----------+              +------+                  +-------+
|          |    Opens     |      |  Waits for QA    |       |
|  Commit  +-------------->  PR  +------------------+ Merge |
|          |              |      |                  |       |
+----------+              +------+                  +-------+

+                     Release Lead Time                     +
+-----------------------------------------------------------+
+                                                           +
                                Release Review Lead time
                          +                                 +
                          +---------------------------------+
                          +                                 +
</pre>

**Release Lead Time:** it measures how much time the first commit in the release request takes to reach the master branch. It's quite common tu assume this metric as *from code to deploy time*.

`Formula: (merged_at  - first_commit_created_at)`

**Release Review Lead Time:** it measures how much time the release review takes, in a lot of organizations it means QA time.

`Formula: (merged_at – opened_at)`

**Relase Size:** it measures the release size in terms of changes it contains.


`Formula: (pull_request_additions + pull_request_deletions)`


## Limitations

It’s an excellent metric to understand whether your team adopted code review as part of their daily activities.
 If the number is considerably high for your context, then you should ask yourself, why are people afraid or avoiding to merge?
 
  If the number is considerably high for your context, then you should ask yourself, why are people afraid or avoiding to merge?
