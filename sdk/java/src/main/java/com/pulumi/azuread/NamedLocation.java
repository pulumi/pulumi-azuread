// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.NamedLocationArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.NamedLocationState;
import com.pulumi.azuread.outputs.NamedLocationCountry;
import com.pulumi.azuread.outputs.NamedLocationIp;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Named Location within Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ConditionalAccess` and `Policy.Read.All`
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Administrator`
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.NamedLocation;
 * import com.pulumi.azuread.NamedLocationArgs;
 * import com.pulumi.azuread.inputs.NamedLocationIpArgs;
 * import com.pulumi.azuread.inputs.NamedLocationCountryArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example_ip = new NamedLocation("example-ip", NamedLocationArgs.builder()        
 *             .displayName("IP Named Location")
 *             .ip(NamedLocationIpArgs.builder()
 *                 .ipRanges(                
 *                     "1.1.1.1/32",
 *                     "2.2.2.2/32")
 *                 .trusted(true)
 *                 .build())
 *             .build());
 * 
 *         var example_country = new NamedLocation("example-country", NamedLocationArgs.builder()        
 *             .displayName("Country Named Location")
 *             .country(NamedLocationCountryArgs.builder()
 *                 .countriesAndRegions(                
 *                     "GB",
 *                     "US")
 *                 .includeUnknownCountriesAndRegions(false)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Named Locations can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/namedLocation:NamedLocation my_location 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuread:index/namedLocation:NamedLocation")
public class NamedLocation extends com.pulumi.resources.CustomResource {
    /**
     * A `country` block as documented below, which configures a country-based named location.
     * 
     */
    @Export(name="country", refs={NamedLocationCountry.class}, tree="[0]")
    private Output</* @Nullable */ NamedLocationCountry> country;

    /**
     * @return A `country` block as documented below, which configures a country-based named location.
     * 
     */
    public Output<Optional<NamedLocationCountry>> country() {
        return Codegen.optional(this.country);
    }
    /**
     * The friendly name for this named location.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The friendly name for this named location.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * An `ip` block as documented below, which configures an IP-based named location.
     * 
     * &gt; Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
     * 
     */
    @Export(name="ip", refs={NamedLocationIp.class}, tree="[0]")
    private Output</* @Nullable */ NamedLocationIp> ip;

    /**
     * @return An `ip` block as documented below, which configures an IP-based named location.
     * 
     * &gt; Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
     * 
     */
    public Output<Optional<NamedLocationIp>> ip() {
        return Codegen.optional(this.ip);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NamedLocation(String name) {
        this(name, NamedLocationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NamedLocation(String name, NamedLocationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NamedLocation(String name, NamedLocationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/namedLocation:NamedLocation", name, args == null ? NamedLocationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NamedLocation(String name, Output<String> id, @Nullable NamedLocationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/namedLocation:NamedLocation", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static NamedLocation get(String name, Output<String> id, @Nullable NamedLocationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NamedLocation(name, id, state, options);
    }
}
