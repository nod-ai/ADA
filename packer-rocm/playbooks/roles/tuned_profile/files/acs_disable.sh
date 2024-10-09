#!/bin/bash
#
# Disable ACS on every device that supports it; offers verification for 'tuned'
#

function force_acs() {
        # default path/function - if not 'verify' mode, *set* ACS/DT mode for devices
        for BDF in $(lspci -d "*:*:*" | awk '{print $1}'); do
                # skip if it doesn't support ACS
                if ! setpci -v -s "${BDF}" ECAP_ACS+0x6.w > /dev/null 2>&1 ; then
                        #echo "${BDF} does not support ACS, skipping"
                        continue
                fi
                echo "Disabling ACS on $(lspci -s "${BDF}")"
                if ! setpci -v -s "${BDF}" ECAP_ACS+0x6.w=0000 ; then
                        echo "Error enabling directTrans ACS on ${BDF}"
                        continue
                fi
                NEW_VAL=$(setpci -v -s "${BDF}" ECAP_ACS+0x6.w | awk '{print $NF}')
                if [ "${NEW_VAL}" != "0000" ]; then
                        echo "Failed to enabling directTrans ACS on ${BDF}"
                        continue
                fi
        done
}

function verify() {
        # if the first/only arg is 'verify', this is called to *check* ACS on all devices
	ACS_ENABLED_CT=$(lspci -vvv | grep -c "ACSCtl.*SrcValid+")
	printf "Verification mode enabled for 'tuned-adm verify'\nFound %s devices with ACS enabled\n" "$ACS_ENABLED_CT"
	# exit with rc matching the number of devices with ACS enabled; > 0 is failure
	exit "$ACS_ENABLED_CT"
}

case $1 in
        verify)
		verify
                ;;
        *)
                force_acs
                ;;
esac
