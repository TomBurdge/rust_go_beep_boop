use std::ffi::CStr;

#[no_mangle]
pub extern "C" fn speak(name: *const libc::c_char) {
    let name_cstr = unsafe { CStr::from_ptr(name) };
    let name = name_cstr.to_str().unwrap();
    println!("beep boop I got your string. here it is: {}!", name);
}

#[cfg(test)]
pub mod test {

    use super::*;
    use std::ffi::CString;

    // This is meant to do the same stuff as the main function in the .go files or other FFI interfaces
    #[test]
    fn simulated_main_function() {
        speak(CString::new("world").unwrap().into_raw());
    }
}
