import { X } from "lucide-react";
import React from "react";
import AsyncSelect from "react-select/async";
import { searchUsers } from "../../api/user.ts";
import {type UserType} from "../../types/user.ts";

interface AddUserToGroupModalProps {
    isOpen: boolean;
    onClose: () => void;
    onSelectUser: (userIds: number[]) => void;
    onSubmit: (e: React.SubmitEvent) => void;
}

function AddUserToGroupModal({
    isOpen,
    onClose,
    onSelectUser,
    onSubmit,
}: AddUserToGroupModalProps) {
    if (!isOpen) return null;

    const loadOptions = (inputValue: string) => new Promise<{value: number, label: string}[]>((resolve, reject) =>
            searchUsers(inputValue).then((users: UserType[]) => {
                const options: {value: number, label: string}[]  = users.map((user) => ({
                    value: user.ID as number,
                    label: user.user_name,
                }));
                resolve(options);
            }).catch(() => {
                reject([]);
            })
    )


    return (
        <>
            {isOpen && (
                <div
                    className="fixed inset-0 bg-(--card) bg-opacity-50 z-40 transition-opacity duration-200"
                    onClick={onClose}
                />
            )}
            {isOpen && (
                <div className="fixed inset-0 z-50 flex items-center justify-center p-4 pointer-events-none">
                    <div
                        className="bg-(--card) rounded-lg border border-[#C5C3C3] p-8 w-full max-w-md shadow-lg pointer-events-auto"
                        onClick={(e) => e.stopPropagation()}
                    >
                        {/* Modal Header */}
                        <div className="flex items-center justify-between mb-6">
                            <h2 className="text-2xl font-bold text-foreground">Add User to Group</h2>
                            <button
                                onClick={onClose}
                                className="p-1 hover:bg-accent rounded-lg transition-colors duration-200"
                            >
                                <X size={24} className="text-muted-foreground" />
                            </button>
                        </div>

                        {/* Form */}
                        <form onSubmit={onSubmit} className="space-y-4">
                            <div>
                                <label htmlFor="user" className="block text-sm font-medium text-foreground mb-2">
                                    Search for users to add *
                                </label>
                                <AsyncSelect
                                    id="user"
                                    name="user"
                                    isMulti
                                    loadOptions={loadOptions}
                                    onChange={(selectedOptions) => {
                                        const userIds = selectedOptions ? selectedOptions.map(option => option.value) : [];
                                        onSelectUser(userIds as number[]);
                                    }}
                                    className="w-full text-foreground placeholder-muted-foreground focus:outline-none focus:ring-2 focus:ring-primary transition-all duration-200"
                                />
                            </div>

                            {/* Form Actions */}
                            <div className="flex gap-3 pt-4">
                                <button
                                    type="button"
                                    onClick={onClose}
                                    className="flex-1 px-4 py-2 bg-accent text-foreground rounded-lg hover:bg-opacity-80 transition-colors duration-200 font-medium"
                                >
                                    Cancel
                                </button>
                                <button
                                    type="submit"
                                    className="flex-1 px-4 py-2 bg-primary text-(--primary) border border-(--primary) rounded-lg hover:bg-opacity-90 transition-all duration-200 font-medium"
                                >
                                    Add Users
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            )}
        </>
    );
}

export default AddUserToGroupModal;
