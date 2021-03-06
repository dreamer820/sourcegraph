import classNames from 'classnames'
import CheckCircleIcon from 'mdi-react/CheckCircleIcon'
import ErrorIcon from 'mdi-react/ErrorIcon'
import FileUploadIcon from 'mdi-react/FileUploadIcon'
import TimerSandIcon from 'mdi-react/TimerSandIcon'
import React, { FunctionComponent } from 'react'

import { LoadingSpinner } from '@sourcegraph/react-loading-spinner'

import { LSIFIndexState, LSIFUploadState } from '../../../graphql-operations'

export interface CodeIntelStateIconProps {
    state: LSIFUploadState | LSIFIndexState
    className?: string
}

export const CodeIntelStateIcon: FunctionComponent<CodeIntelStateIconProps> = ({ state, className }) =>
    state === LSIFUploadState.UPLOADING ? (
        <FileUploadIcon className={classNames('redesign-d-none', className)} />
    ) : state === LSIFUploadState.QUEUED || state === LSIFIndexState.QUEUED ? (
        <TimerSandIcon className={classNames('redesign-d-none', className)} />
    ) : state === LSIFUploadState.PROCESSING || state === LSIFIndexState.PROCESSING ? (
        <LoadingSpinner className={classNames('redesign-d-none', className)} />
    ) : state === LSIFUploadState.COMPLETED || state === LSIFIndexState.COMPLETED ? (
        <CheckCircleIcon className={classNames('redesign-d-none text-success', className)} />
    ) : state === LSIFUploadState.ERRORED || state === LSIFIndexState.ERRORED ? (
        <ErrorIcon className={classNames('redesign-d-none text-danger', className)} />
    ) : (
        <></>
    )
