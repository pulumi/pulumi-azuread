// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SynchronizationJobSchedule {
    /**
     * @return Date and time when this job will expire, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     * 
     */
    private @Nullable String expiration;
    /**
     * @return The interval between synchronization iterations ISO8601. E.g. PT40M run every 40 minutes.
     * 
     */
    private @Nullable String interval;
    /**
     * @return State of the job.
     * 
     */
    private @Nullable String state;

    private SynchronizationJobSchedule() {}
    /**
     * @return Date and time when this job will expire, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     * 
     */
    public Optional<String> expiration() {
        return Optional.ofNullable(this.expiration);
    }
    /**
     * @return The interval between synchronization iterations ISO8601. E.g. PT40M run every 40 minutes.
     * 
     */
    public Optional<String> interval() {
        return Optional.ofNullable(this.interval);
    }
    /**
     * @return State of the job.
     * 
     */
    public Optional<String> state() {
        return Optional.ofNullable(this.state);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SynchronizationJobSchedule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String expiration;
        private @Nullable String interval;
        private @Nullable String state;
        public Builder() {}
        public Builder(SynchronizationJobSchedule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.expiration = defaults.expiration;
    	      this.interval = defaults.interval;
    	      this.state = defaults.state;
        }

        @CustomType.Setter
        public Builder expiration(@Nullable String expiration) {
            this.expiration = expiration;
            return this;
        }
        @CustomType.Setter
        public Builder interval(@Nullable String interval) {
            this.interval = interval;
            return this;
        }
        @CustomType.Setter
        public Builder state(@Nullable String state) {
            this.state = state;
            return this;
        }
        public SynchronizationJobSchedule build() {
            final var _resultValue = new SynchronizationJobSchedule();
            _resultValue.expiration = expiration;
            _resultValue.interval = interval;
            _resultValue.state = state;
            return _resultValue;
        }
    }
}
