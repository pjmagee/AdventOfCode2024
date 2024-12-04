cmake_minimum_required(VERSION 3.30)

if (DEFINED ENV{VCPKG_ROOT} AND NOT DEFINED CMAKE_TOOLCHAIN_FILE)
    set(CMAKE_TOOLCHAIN_FILE "$ENV{VCPKG_ROOT}/scripts/buildsystems/vcpkg.cmake"
            CACHE STRING "")
endif ()

project(advent)

set(CMAKE_CXX_STANDARD 26)

find_package(nlohmann_json CONFIG REQUIRED)

set(nlohmann-json_IMPLICIT_CONVERSIONS OFF)

add_executable(advent
        Day.h
        DayFactory.h
        DayFactory.cpp
        Day1.h
        Day1.cpp
        Day2.h
        Day2.cpp
        Day3.h
        Day3.cpp
        main.cpp
)

if (MINGW)
    set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -mconsole")
    set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -static")
endif ()

target_compile_options(advent PRIVATE -fno-rtti)
target_link_libraries(advent PRIVATE nlohmann_json::nlohmann_json)

add_custom_target(clean-all
        COMMAND ${CMAKE_COMMAND} -P ${CMAKE_CURRENT_SOURCE_DIR}/clean-all.cmake
        COMMENT "Cleaning all")