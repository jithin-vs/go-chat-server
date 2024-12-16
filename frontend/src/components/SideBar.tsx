"use client";
import { BsSearch } from "react-icons/bs";
import SideUser from "./chatLayout/sideUser";
import React, { useEffect, useState } from "react";
import axios from "axios";
import { PORT } from "@/config";
import Search from "./search";
import { acessCreate } from "@/apis";
import { useAuth } from "@/context/AuthContext";

export default function SideBar() {
  const [search, setSearch] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [searchResults, setSearchResults] = useState([]);
  const { user } = useAuth();
  let currentUserId = user?.id;

  const handleSearch = async (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    setSearch(e.target.value);
  };

  const handleClick = async (e: any) => {
    const body = {
      userId: e._id,
      senderId:currentUserId
    }
    await acessCreate(body)
    setSearch("")
  }

  const searchUsers = async (query: string) => {
    axios
      .get(`${PORT}/api/user/search`, {
        params: {
          query: query, // Pass your search keyword
        },
        headers: {
          "Content-Type": "application/json",
        },
      })
      .then((response) => {
        if (response.status === 200) {
          console.log(
            "Search results fetched successfully:",
            response.data.results
          );

          // Process the search results
          const searchResults = response.data.results;

          if (searchResults.length > 0) {
            setSearchResults(searchResults); // Update state to display results
          } else {
            console.warn("No matching results found");
          }
        } else {
          console.log("Unexpected status while searching:", response);
        }
      })
      .catch((error) => {
        console.error("Error fetching search results:", error);
        // Handle errors (e.g., show error to user)
        // setErrors([error.response?.data?.error || "Unknown error occurred"]);
      });
  };

  useEffect(() => {
    const searchChange = async () => {
      setIsLoading(true);
      searchUsers(search);
      setIsLoading(false);
    };
    searchChange();
  }, [search]);
  return (
    <div className="w-1/4 bg-white border-r border-gray-300">
      {/* <!-- Sidebar Header --> */}
      <header className="p-4 border-b border-gray-300 flex justify-between items-center bg-indigo-600 text-white">
        <h1 className="text-2xl font-semibold">Chat Web</h1>
        <div className="relative">
          <form onSubmit={(e) => e.preventDefault()}>
            <input
              onChange={handleSearch}
              className="w-[99.5%] bg-[#f6f6f6] text-[#111b21] tracking-wider pl-9 py-[8px] rounded-[9px] outline-0"
              type="text"
              name="search"
              placeholder="Search"
            />
          </form>

          <div className="absolute top-[36px] left-[27px]">
            <BsSearch style={{ color: "#c4c4c5" }} />
          </div>
          <Search searchResults={searchResults} isLoading={isLoading} handleClick={handleClick} search={search}/>
          <button id="menuButton" className="focus:outline-none">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-5 w-5 text-gray-100"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
              <path d="M2 10a2 2 0 012-2h12a2 2 0 012 2 2 2 0 01-2 2H4a2 2 0 01-2-2z" />
            </svg>
          </button>
          {/* <!-- Menu Dropdown --> */}
          <div
            id="menuDropdown"
            className="absolute right-0 mt-2 w-48 bg-white border border-gray-300 rounded-md shadow-lg hidden"
          >
            <ul className="py-2 px-3">
              <li>
                <a
                  href="#"
                  className="block px-4 py-2 text-gray-800 hover:text-gray-400"
                >
                  Option 1
                </a>
              </li>
              <li>
                <a
                  href="#"
                  className="block px-4 py-2 text-gray-800 hover:text-gray-400"
                >
                  Option 2
                </a>
              </li>
              {/* <!-- Add more menu options here --> */}
            </ul>
          </div>
        </div>
      </header>

      {/* <!-- Contact List --> */}
      <div className="overflow-y-auto h-screen p-3 mb-9 pb-20">
        <SideUser />
      </div>
    </div>
  );
}
