# dx is a tool to generate cfg2 xml from input data.

## usage:

### -help for help

$ dx -h

mp: plag-telemetry action: ponf

generates plat-telemetry cfg2 for ponf

$ dx -p plat-telemetry -a ponf data.txt output.xml

## input csv format

First line should be the metadata which describes column name with comma delimiter.
Afterwards would be the data payload.

e.g

```csv
osgroup,osname,osversion,ip_ver,tcp_ttl,tcp_olen,tcp_mss,tcp_win,tcp_wscale,tcp_ipoptions,tcp_quirks,pclass
unix,Linux,3.19 and newer,*,68,0,*,mss*20,10,2-4-8-1-3,6,0
unix,Linux,3.10 and newer,*,64,0,*,mss*20,10,2-4-8-1-3,6,0
unix,Linux,3.12 and newer,*,60,0,*,mss*20,10,2-4-8-1-3,6,0
```

### generated cfg2 xml format as follows:

```javascript
<application
    xmlns="urn:com:linkedin:ns:configuration:source:1.0"
    xmlns:w="urn:com:linkedin:ns:configuration:wildcard:1.0"
>
    <configuration-source>
        <property
            name="ponfMetaData"
            value="osgroup,osname,osversion,ip_ver,tcp_ttl,tcp_olen,tcp_mss,tcp_win,tcp_wscale,tcp_ipoptions,tcp_quirks,pclass"
        ></property>
        <property name="ponfFeatures">
            <set>
                <value>unix,Linux,3.19 and newer,*,68,0,*,mss*20,10,2-4-8-1-3,6,0</value>
                <value>unix,Linux,3.10 and newer,*,64,0,*,mss*20,10,2-4-8-1-3,6,0</value>
                <value>unix,Linux,3.12 and newer,*,60,0,*,mss*20,10,2-4-8-1-3,6,0</value>
            </set>
        </property>
    </configuration-source>
</application>
```
