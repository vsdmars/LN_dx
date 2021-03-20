package main

/*
dx is a tool to generate cfg2 xml from ponf signatures cvs file.

usage:
dx input_csv_file_name output_xml_file_name

generated cfg2 xml format as follows:

	<application xmlns="urn:com:linkedin:ns:configuration:source:1.0" xmlns:w="urn:com:linkedin:ns:configuration:wildcard:1.0">
	   <configuration-source>
		 <property name="metaData" value="col1,col2,col3"></property>
		 <property>
		   <set>
			 <value>val1,val2,val3</value>
			 <value>val4,val5,val6</value>
		   </set>
		 </property>
	   </configuration-source>
	 </application>


Todo:
1. CLI
2. input can indicates property type
name|value|data
name|type|data

*/
