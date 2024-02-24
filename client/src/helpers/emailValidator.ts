// Basic email validation using a regular expression
export const isEmailValid = (email: string): boolean => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
};

//  interface ValidationResponse {
//     isValid: boolean,
//     error: string
//  }

// export const validateRegistrationForm = (email: string, password: string): boolean => {
//     let isValid: boolean = true;
//     switch (true) {
//         case !email || !password:
            
//           break;
//         case !isEmailValid(email):
//           setError("Enter valid email");
//           break;
//         case password.length <= 3:
//           setError("Password must be longer than 3 characters");
//           break;
//         default:
//           setError('');
//           break;
//       }
//     return true
// }