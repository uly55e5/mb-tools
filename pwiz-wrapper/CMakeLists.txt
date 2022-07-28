CMAKE_MINIMUM_REQUIRED(VERSION 3.22)
PROJECT(pwiz_wrapper)
set(CMAKE_VERBOSE_MAKEFILE ON)
SET(VERBOSE ON)
SET(PWIZ_DIR "" CACHE PATH "Path to Protein wizard source code")
IF(PWIZ_DIR STREQUAL "")
  MESSAGE(WARNING "No path for ProteinWizard source is given. Set PWIZ_DIR")
ENDIF()

  EXECUTE_PROCESS(COMMAND ./quickbuild.sh pwiz/data/msdata WORKING_DIRECTORY ${PWIZ_DIR})
  FILE(GLOB_RECURSE PWIZ_OBJ "${PWIZ_DIR}/build-linux-x86_64/*/libpwiz_data_msdata.a")
  INCLUDE_DIRECTORIES(. ${PWIZ_DIR})

SET(CMAKE_ARCHIVE_OUTPUT_DIRECTORY_RELEASE ${CMAKE_SOURCE_DIR}/lib)

ADD_LIBRARY(pwiz_wrapper STATIC
               cpwiz.cpp
               cpwiz.h
               ${PWIZ_DIR}/build-linux-x86_64/pwiz/data/msdata/gcc-9/release/link-static/runtime-link-static/threading-multi/)
