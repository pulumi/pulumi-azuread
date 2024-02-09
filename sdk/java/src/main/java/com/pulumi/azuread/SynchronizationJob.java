// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.SynchronizationJobArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.SynchronizationJobState;
import com.pulumi.azuread.outputs.SynchronizationJobSchedule;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a synchronization job associated with a service principal (enterprise application) within Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`
 * 
 * ## Example Usage
 * 
 * *Basic example*
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.inputs.GetApplicationTemplateArgs;
 * import com.pulumi.azuread.Application;
 * import com.pulumi.azuread.ApplicationArgs;
 * import com.pulumi.azuread.inputs.ApplicationFeatureTagArgs;
 * import com.pulumi.azuread.ServicePrincipal;
 * import com.pulumi.azuread.ServicePrincipalArgs;
 * import com.pulumi.azuread.SynchronizationSecret;
 * import com.pulumi.azuread.SynchronizationSecretArgs;
 * import com.pulumi.azuread.inputs.SynchronizationSecretCredentialArgs;
 * import com.pulumi.azuread.SynchronizationJob;
 * import com.pulumi.azuread.SynchronizationJobArgs;
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
 *         final var exampleApplicationTemplate = AzureadFunctions.getApplicationTemplate(GetApplicationTemplateArgs.builder()
 *             .displayName(&#34;Azure Databricks SCIM Provisioning Connector&#34;)
 *             .build());
 * 
 *         var exampleApplication = new Application(&#34;exampleApplication&#34;, ApplicationArgs.builder()        
 *             .displayName(&#34;example&#34;)
 *             .templateId(exampleApplicationTemplate.applyValue(getApplicationTemplateResult -&gt; getApplicationTemplateResult.templateId()))
 *             .featureTags(ApplicationFeatureTagArgs.builder()
 *                 .enterprise(true)
 *                 .gallery(true)
 *                 .build())
 *             .build());
 * 
 *         var exampleServicePrincipal = new ServicePrincipal(&#34;exampleServicePrincipal&#34;, ServicePrincipalArgs.builder()        
 *             .applicationId(exampleApplication.applicationId())
 *             .useExisting(true)
 *             .build());
 * 
 *         var exampleSynchronizationSecret = new SynchronizationSecret(&#34;exampleSynchronizationSecret&#34;, SynchronizationSecretArgs.builder()        
 *             .servicePrincipalId(exampleServicePrincipal.id())
 *             .credentials(            
 *                 SynchronizationSecretCredentialArgs.builder()
 *                     .key(&#34;BaseAddress&#34;)
 *                     .value(&#34;https://adb-example.azuredatabricks.net/api/2.0/preview/scim&#34;)
 *                     .build(),
 *                 SynchronizationSecretCredentialArgs.builder()
 *                     .key(&#34;SecretToken&#34;)
 *                     .value(&#34;some-token&#34;)
 *                     .build())
 *             .build());
 * 
 *         var exampleSynchronizationJob = new SynchronizationJob(&#34;exampleSynchronizationJob&#34;, SynchronizationJobArgs.builder()        
 *             .servicePrincipalId(exampleServicePrincipal.id())
 *             .templateId(&#34;dataBricks&#34;)
 *             .enabled(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Synchronization jobs can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/synchronizationJob:SynchronizationJob example 00000000-0000-0000-0000-000000000000/job/dataBricks.f5532fc709734b1a90e8a1fa9fd03a82.8442fd39-2183-419c-8732-74b6ce866bd5
 * ```
 * 
 *  -&gt; This ID format is unique to Terraform and is composed of the Service Principal Object ID and the ID of the Synchronization Job Id in the format `{servicePrincipalId}/job/{jobId}`.
 * 
 */
@ResourceType(type="azuread:index/synchronizationJob:SynchronizationJob")
public class SynchronizationJob extends com.pulumi.resources.CustomResource {
    /**
     * Whether or not the provisioning job is enabled. Default state is `true`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Whether or not the provisioning job is enabled. Default state is `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * A `schedule` list as documented below.
     * 
     */
    @Export(name="schedules", refs={List.class,SynchronizationJobSchedule.class}, tree="[0,1]")
    private Output<List<SynchronizationJobSchedule>> schedules;

    /**
     * @return A `schedule` list as documented below.
     * 
     */
    public Output<List<SynchronizationJobSchedule>> schedules() {
        return this.schedules;
    }
    /**
     * The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
     * 
     */
    @Export(name="servicePrincipalId", refs={String.class}, tree="[0]")
    private Output<String> servicePrincipalId;

    /**
     * @return The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> servicePrincipalId() {
        return this.servicePrincipalId;
    }
    /**
     * Identifier of the synchronization template this job is based on.
     * 
     */
    @Export(name="templateId", refs={String.class}, tree="[0]")
    private Output<String> templateId;

    /**
     * @return Identifier of the synchronization template this job is based on.
     * 
     */
    public Output<String> templateId() {
        return this.templateId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SynchronizationJob(String name) {
        this(name, SynchronizationJobArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SynchronizationJob(String name, SynchronizationJobArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SynchronizationJob(String name, SynchronizationJobArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/synchronizationJob:SynchronizationJob", name, args == null ? SynchronizationJobArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SynchronizationJob(String name, Output<String> id, @Nullable SynchronizationJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/synchronizationJob:SynchronizationJob", name, state, makeResourceOptions(options, id));
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
    public static SynchronizationJob get(String name, Output<String> id, @Nullable SynchronizationJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SynchronizationJob(name, id, state, options);
    }
}
