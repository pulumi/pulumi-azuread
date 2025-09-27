# Install Pulumi and plugins required at build time.
install_plugins: .make/install_plugins
.make/install_plugins: export PULUMI_HOME := $(WORKING_DIR)/.pulumi
.make/install_plugins:
	pulumi plugin install resource time 0.0.15
	pulumi plugin install resource std 1.4.0
	pulumi plugin install resource azure 6.3.0
	pulumi plugin install resource random 4.16.0
	pulumi plugin install converter terraform 1.0.15
	@touch $@
.PHONY: install_plugins
