use std::{cell::RefCell, rc::Rc};

use codec::frame::{Frame, OperationCode};
use protocol::rpc::header::ErrorCode;
use slog::Logger;
use store::ElasticStore;

use crate::workspace::stream_manager::StreamManager;

use super::{
    append::Append, describe_range::DescribeRange, fetch::Fetch, heartbeat::Heartbeat, ping::Ping,
    seal_range::SealRange,
};

#[derive(Debug)]
pub(crate) enum Command<'a> {
    Append(Append<'a>),
    DescribeRange(DescribeRange<'a>),
    Fetch(Fetch<'a>),
    SealRange(SealRange<'a>),
    Ping(Ping<'a>),
    Heartbeat(Heartbeat<'a>),
}

impl<'a> Command<'a> {
    pub fn from_frame(logger: Logger, frame: &Frame) -> Result<Command, ErrorCode> {
        match frame.operation_code {
            OperationCode::DescribeRanges => Ok(Command::DescribeRange(
                DescribeRange::parse_frame(logger.clone(), frame)?,
            )),

            OperationCode::Unknown => todo!(),

            OperationCode::Ping => Ok(Command::Ping(Ping::new(logger.clone(), frame))),

            OperationCode::GoAway => todo!(),

            OperationCode::Heartbeat => Ok(Command::Heartbeat(Heartbeat::parse_frame(
                logger.clone(),
                frame,
            )?)),

            OperationCode::AllocateId => unreachable!(),

            OperationCode::Append => {
                Ok(Command::Append(Append::parse_frame(logger.clone(), frame)?))
            }

            OperationCode::Fetch => Ok(Command::Fetch(Fetch::parse_frame(logger.clone(), frame)?)),

            OperationCode::ListRanges => todo!(),

            OperationCode::SealRanges => Ok(Command::SealRange(SealRange::parse_frame(
                logger.clone(),
                frame,
            )?)),

            OperationCode::SyncRanges => todo!(),
            OperationCode::CreateStreams => todo!(),
            OperationCode::DeleteStreams => todo!(),
            OperationCode::UpdateStreams => todo!(),
            OperationCode::DescribeStreams => todo!(),
            OperationCode::TrimStreams => todo!(),
            OperationCode::ReportMetrics => todo!(),
        }
    }

    pub(crate) async fn apply(
        &self,
        store: Rc<ElasticStore>,
        stream_manager: Rc<RefCell<StreamManager>>,
        response: &mut Frame,
    ) {
        match self {
            Command::Append(cmd) => cmd.apply(store, stream_manager, response).await,
            Command::DescribeRange(cmd) => cmd.apply(store, stream_manager, response).await,
            Command::Fetch(cmd) => cmd.apply(store, stream_manager, response).await,
            Command::Heartbeat(cmd) => cmd.apply(store, stream_manager, response).await,
            Command::Ping(cmd) => cmd.apply(store, stream_manager, response).await,
            Command::SealRange(cmd) => cmd.apply(store, stream_manager, response).await,
        }
    }
}
