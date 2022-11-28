compile kotlin
    kotlinc hello.kt -include-runtime -d hello.jar
need package
    ops pkg get eyberg/java:19.6.625
create with
    ops image create -c config.json -i nanos_kotlin_bm --package eyberg/java:19.6.625