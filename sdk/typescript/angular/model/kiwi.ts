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
 * A Kiwi (Kowabunga Inner Wan Interface) provides edge-network services..
 */
export interface Kiwi { 
    /**
     * The Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. ID (auto-generated).
     */
    id?: string;
    /**
     * The Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. name.
     */
    name: string;
    /**
     * The Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. description.
     */
    description?: string;
    /**
     * a list of existing remote agents managing the network gateway.
     */
    agents?: Array<string>;
}

