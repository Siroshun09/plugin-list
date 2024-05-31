plugins {
    java
    id("io.github.goooler.shadow") version "8.1.7"
}

group = "dev.siroshun.pluginlistcollector"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
    maven(url = "https://repo.papermc.io/repository/maven-public/")

    mavenLocal()
    maven(url = " https://oss.sonatype.org/content/repositories/snapshots/")
}

dependencies {
    compileOnly("io.papermc.paper:paper-api:1.20.6-R0.1-SNAPSHOT")

    implementation(files("libs/codec4j-api-0.0.1-SNAPSHOT.jar", "libs/codec4j-format-gson-0.0.1-SNAPSHOT.jar"))

    testImplementation(platform("org.junit:junit-bom:5.10.2"))
    testImplementation("org.junit.jupiter:junit-jupiter:5.10.2")
    testImplementation("com.google.code.gson:gson:2.11.0")
    testRuntimeOnly("org.junit.platform:junit-platform-launcher")
}

java {
    sourceCompatibility = JavaVersion.VERSION_21
    targetCompatibility = JavaVersion.VERSION_21
}

tasks {
    build {
        dependsOn(shadowJar)
    }

    compileJava {
        options.encoding = Charsets.UTF_8.name()
        options.release.set(21)
    }

    processResources {
        filesMatching(listOf("paper-plugin.yml")) {
            expand("projectVersion" to version)
        }
    }

    test {
        useJUnitPlatform()
    }

    jar {
        archiveFileName = "${project.name}-${project.version}-original.jar"
    }

    shadowJar {
        minimize()
        archiveFileName = "PluginListCollector-${project.version}.jar"
    }
}
