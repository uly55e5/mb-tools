CMAKE_MINIMUM_REQUIRED(VERSION 3.22)
PROJECT(pwiz_wrapper)
set(CMAKE_VERBOSE_MAKEFILE ON)
SET(VERBOSE ON)
SET(PWIZ_DIR "" CACHE PATH "Path to Protein wizard source code")
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
IF(PWIZ_DIR STREQUAL "")
  MESSAGE(FATAL "No path for ProteinWizard source is given. Set PWIZ_DIR")
ENDIF()
IF(NOT EXISTS ${CMAKE_SOURCE_DIR}/lib/libpwiz_all.a)
  EXECUTE_PROCESS(COMMAND ./quickbuild.sh pwiz/data/msdata --clean WORKING_DIRECTORY ${PWIZ_DIR})
  EXECUTE_PROCESS(COMMAND ./quickbuild.sh pwiz/data/msdata WORKING_DIRECTORY ${PWIZ_DIR})
  FILE(GLOB_RECURSE PWIZ_OBJ "${PWIZ_DIR}/**/*.o")
  EXECUTE_PROCESS(COMMAND ar -rcs ${CMAKE_SOURCE_DIR}/lib/libpwiz_all.a ${PWIZ_OBJ})
ENDIF()
INCLUDE_DIRECTORIES(. ${PWIZ_DIR})

SET(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR}/lib)
set(CMAKE_DEBUG_POSTFIX d)
ADD_LIBRARY(pwiz_wrapper STATIC
               cpwiz.cpp
               cpwiz.h)

ADD_EXECUTABLE(test
               main-test.cpp)

SET(CMAKE_EXE_LINKER_FLAGS "-no-pie -lm -ldl -lstdc++ -pthread")
TARGET_LINK_LIBRARIES(test pwiz_wrapper ${CMAKE_SOURCE_DIR}/lib/libpwiz_all.a)