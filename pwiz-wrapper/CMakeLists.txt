CMAKE_MINIMUM_REQUIRED(VERSION 3.22)
PROJECT(pwiz_wrapper)

SET(CMAKE_CXX_STANDARD 14)

SET(PWIZ_LIB pwizmsdata.a)

SET(PWIZ_DIR "" CACHE PATH "Path to Protein wizard source code")
IF(PWIZ_DIR STREQUAL "")
  MESSAGE(WARNING "No path for ProteinWizard source is given. Set PWIZ_DIR")
ENDIF()

IF(NOT EXISTS ${PWIZ_LIB})
  EXECUTE_PROCESS(COMMAND ./quickbuild.sh pwiz/data/msdata WORKING_DIRECTORY ${PWIZ_DIR})
  FILE(GLOB_RECURSE PWIZ_OBJ ${PWIZ_DIR}/build-linux-x86_64/*.o)
  EXECUTE_PROCESS(COMMAND ar rvs ${PWIZ_LIB} ${PWIZ_OBJ}
                WORKING_DIRECTORY .)
ENDIF()
IF(NOT EXISTS ${PWIZ_LIB})
  MESSAGE(FATAL_ERROR "Could not build the Proteinwizard library")
ENDIF()

INCLUDE_DIRECTORIES(. ${PWIZ_DIR})

SET(CMAKE_ARCHIVE_OUTPUT_DIRECTORY_RELEASE ${CMAKE_SOURCE_DIR}/lib)

ADD_LIBRARY(pwiz_wrapper
               cpwiz.cpp
               cpwiz.h)

LINK_LIBRARIES(pwiz_wrapper ${PWIZ_LIB})
