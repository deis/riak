
|![](https://upload.wikimedia.org/wikipedia/commons/thumb/1/17/Warning.svg/156px-Warning.svg.png) | Deis Workflow is no longer maintained.<br />Please [read the announcement](https://deis.com/blog/2017/deis-workflow-final-release/) for more detail. |
|---:|---|
| 09/07/2017 | Deis Workflow [v2.18][] final release before entering maintenance mode |
| 03/01/2018 | End of Workflow maintenance: critical patches no longer merged |
| | [Hephy](https://github.com/teamhephy/workflow) is a fork of Workflow that is actively developed and accepts code contributions. |

# Riak

[![Build Status](https://travis-ci.org/deis/riak.svg?branch=master)](https://travis-ci.org/deis/riak)

This is a project for running Riak in a kubernetes cluster.

## Usage

If you want to experiment with running this component in kubernetes, try this:

```
$ make # builds and pushes the image
$ make deploy
```

At this point, 3 nodes should come up. However, Riak provides a multi-phased approach to cluster
administration that enables you to stage and review cluster-level changes prior to committing them.
Our cluster is in a "staged" state which means that the cluster nodes are ready to speak with each
other, but the transaction hasn't yet been completed. To commit these changes so the cluster will
start replicating data, run

```
$ kubectl --namespace=deis exec deis-riak-bootstrap riak-admin cluster plan
$ kubectl --namespace=deis exec deis-riak-bootstrap riak-admin cluster commit
```

## Debugging

To see what's going on in your Riak cluster, feel free to poke and prod it by running

```
$ kubectl --namespace=deis exec deis-riak-bootstrap riak-admin member-status
================================= Membership ==================================
Status     Ring    Pending    Node
-------------------------------------------------------------------------------
valid      32.8%      --      'riak@10.246.1.34'
valid      32.8%      --      'riak@10.246.1.35'
valid      34.4%      --      'riak@10.246.1.36'
-------------------------------------------------------------------------------
Valid:3 / Leaving:0 / Exiting:0 / Joining:0 / Down:0
```

Follow the [riak-admin](http://docs.basho.com/riak/latest/ops/running/tools/riak-admin/)
documentation for more info on managing your cluster.

## Scaling

You should have 3 replicas of riak available after deploying initially. To scale up, you can do so
with

```
$ kubectl --namespace=deis scale rc deis-riak --replicas=10
```

Wait for all the replicas to come up, then

```
$ kubectl --namespace=deis exec $DEIS_RIAK_POD_NAME riak-admin cluster plan
```

Review the changes are fit to your desire, then

```
$ kubectl --namespace=deis exec $DEIS_RIAK_POD_NAME riak-admin cluster commit
```

To commit those changes.

## RootFS

All files that are to be packaged into the container should be written
to the `rootfs/` folder.

## Extended Testing

Along with unit tests, Deis values functional and integration testing.
These tests should go in the `_tests` folder.

[v2.18]: https://github.com/deis/workflow/releases/tag/v2.18.0
