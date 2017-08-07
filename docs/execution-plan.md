# Execution Plan #

## Summary
This is a living doc on our current execution strategy.

## Plan on architecture
The client side needs to be lean and clean, with no dependencies. The client will be done in golang. Because of the lack of ML and NLP libraries, the client side will only execute well-structured plans, with some tolerence on human edits, but will not do machine learning.

The server side will collect the swagger.json, does NLP analysis on it, and generate the @meqa tags needed for client to execute. 

The testplan.yml can either be generated by the client or the server.

* Client
    * we can have something fully functioning without the server piece. With just people editing the swagger.json.
    * Not sharing test plans between client and server means we don't have a potential breaking point where the two may be out of sync in versions.
    * Our customers can have better performance and flexibility when generating test plans.
    * Can golang client automatically update itself?
* Server
    * We will have more flexibility in the kind of technology we can leverage. For instance, we can leverage spacy.io during plan generation.
    * When the user modifies and overrides some parameters, we need logic to merge the diffs and learn from the user behavior for future test generation.
    * We can more quickly roll out enhancements - when we employ some new algorithms on the server side to generate better plans, all our clients can immediately benefit. We can tell our customers that if they try generate a new plan, the diff will tell them what new tests are added, and they can decide whether to use the new plan or just the old one.
    * The above plan generation on server will be good for SaaS business.

## Language for Test Plan Generation
The language decision depends on whether there is more code overlap between this and the @meqa tag generation, or the golang test-plan execution. 

* @meqa generation
    * Needs to parse swagger.json.
    * Needs to use NLP.
* Test plan generation
    * Needs to parse swagger.json.
    * Needs to have understanding of the test plan structure.
    * Needs to do swagger structure traversal to generate DAG.
    * Needs to use DAG to generate test plan.
    * Needs to understand parameter generation strategies, and understand how to merge newer plans with the customers' modifications on the client side.
    * It seems any NLP related task can be done through @meqa generation step. If needed we can always generate more tags to guide the following steps.
* Test plan execution
    * Needs to parse swagger.json.
    * Needs to have understanding of the test plan structure.
    * Needs to execute test plan.
    * Needs to modify parameters through trial-and-error (future)

Decision

* Test plan generation is done on the server side in golang.
* This gives us the most flexibility in the beginning.
    * We have best leverage of our code base.
    * We have great flexibility when doing customer trials.

## Server

The server has the following pieces

* swagger.json tagging - generating @meqa tags. This is the lowest priority and we can manually generate the tags for now.
* Generating test plan.
* REST server - account management, communicate with the golang client, support UI.