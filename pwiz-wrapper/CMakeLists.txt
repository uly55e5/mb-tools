CMAKE_MINIMUM_REQUIRED(VERSION 3.22)
PROJECT(pwiz_wrapper)
set(CMAKE_VERBOSE_MAKEFILE ON)
SET(VERBOSE ON)
SET(PREFIX_3rdParty "../3rdparty")
SET(PWIZ_DIR "${CMAKE_SOURCE_DIR}/3rdparty/pwiz" )
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

INCLUDE("ExternalProject")

ExternalProject_Add("pwiz"
                    SOURCE_DIR "${PWIZ_DIR}"
                    GIT_REPOSITORY "https://github.com/ProteoWizard/pwiz.git"
                    GIT_TAG 8c29d2e30b410d56d7673100373639936630d2cd
                    GIT_SHALLOW ON
                    GIT_PROGRESS ON
                    CONFIGURE_COMMAND ""
                    BUILD_COMMAND ./quickbuild.sh pwiz/data/msdata
                    COMMAND find . -name "*.o" | xargs ar -rcs ${CMAKE_SOURCE_DIR}/lib/libpwiz_all.a
                    BUILD_IN_SOURCE ON
                    BUILD_BYPRODUCTS ${CMAKE_SOURCE_DIR}/lib/libpwiz_all.a
                    INSTALL_COMMAND ""
                    )

INCLUDE_DIRECTORIES(. ${PWIZ_DIR})

SET(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR}/lib)
set(CMAKE_DEBUG_POSTFIX d)
ADD_LIBRARY(pwiz_wrapper STATIC
        cpwiz.cpp
        cpwiz.h)

ADD_DEPENDENCIES(pwiz_wrapper pwiz)
ADD_EXECUTABLE(test
        main-test.cpp)

SET(CMAKE_EXE_LINKER_FLAGS "-no-pie -lm -ldl -lstdc++ -pthread")
TARGET_LINK_LIBRARIES(test pwiz_wrapper ${CMAKE_SOURCE_DIR}/lib/libpwiz_all.a)