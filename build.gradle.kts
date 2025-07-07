plugins {
    id("java")
    id("com.google.protobuf") version "0.9.5"
}

group = "week11"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation("com.google.protobuf:protobuf-java:4.30.2")
    implementation("io.grpc:grpc-netty-shaded:1.51.0")
    implementation("io.grpc:grpc-protobuf:1.51.0")
    implementation("io.grpc:grpc-stub:1.51.0")
    implementation("javax.annotation:javax.annotation-api:1.3.2")

    testImplementation(platform("org.junit:junit-bom:5.10.0"))
    testImplementation("org.junit.jupiter:junit-jupiter")
}

protobuf {
    plugins {
        create("grpc") {
            artifact = "io.grpc:protoc-gen-grpc-java:1.51.0"
        }
    }
    generateProtoTasks {
        all().forEach {
            it.plugins {
                create("grpc")
            }
        }
    }
}

tasks.test {
    useJUnitPlatform()
}
