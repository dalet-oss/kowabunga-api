/**
 * Kowabunga
 * Kvm Orchestrator With A BUNch of Goods Added
 *
 * The version of the OpenAPI document: 0.9.1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * The global project resource quotas/usage (0 for unlimited).
 */
export interface ProjectResources { 
    /**
     * the maximum total number of vCPUs allowed to be consumed by sum of all instances.
     */
    vcpus?: number;
    /**
     * the maximum total memory (in bytes) allowed to be consumed by sum of all instances.
     */
    memory?: number;
    /**
     * the maximum total disk capacity allowed to be consumed by sum of all instances.
     */
    storage?: number;
    /**
     * the maximum number of instances allowed to be spawned.
     */
    instances?: number;
}
