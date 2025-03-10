/**
 * Kowabunga API documentation
 *
 * Contact: maintainers@kowabunga.cloud
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * A global project resource quotas/usage (0 for unlimited).
 */
export interface ProjectResources { 
    /**
     * The maximum total number of vCPUs allowed to be consumed by sum of all instances.
     */
    vcpus?: number;
    /**
     * The maximum total memory (in bytes) allowed to be consumed by sum of all instances.
     */
    memory?: number;
    /**
     * The maximum total disk capacity allowed to be consumed by sum of all instances.
     */
    storage?: number;
    /**
     * The maximum number of instances allowed to be spawned.
     */
    instances?: number;
}

